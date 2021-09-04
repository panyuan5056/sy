// Copyright 2016 - 2020 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to
// and read from XLSX / XLSM / XLTM files. Supports reading and writing
// spreadsheet documents generated by Microsoft Exce™ 2007 and later. Supports
// complex components by high compatibility, and provided streaming API for
// generating or reading data from a worksheet with huge amounts of data. This
// library needs Go version 1.10 or later.

package excelize

import (
	"fmt"
	"strings"
)

// MergeCell provides a function to merge cells by given coordinate area and
// sheet name. For example create a merged cell of D3:E9 on Sheet1:
//
//    err := f.MergeCell("Sheet1", "D3", "E9")
//
// If you create a merged cell that overlaps with another existing merged cell,
// those merged cells that already exist will be removed.
//
//                 B1(x1,y1)      D1(x2,y1)
//               +------------------------+
//               |                        |
//     A4(x3,y3) |    C4(x4,y3)           |
//    +------------------------+          |
//    |          |             |          |
//    |          |B5(x1,y2)    | D5(x2,y2)|
//    |          +------------------------+
//    |                        |
//    |A8(x3,y4)      C8(x4,y4)|
//    +------------------------+
//
func (f *File) MergeCell(sheet, hcell, vcell string) error {
	rect1, err := f.areaRefToCoordinates(hcell + ":" + vcell)
	if err != nil {
		return err
	}
	// Correct the coordinate area, such correct C1:B3 to B1:C3.
	_ = sortCoordinates(rect1)

	hcell, _ = CoordinatesToCellName(rect1[0], rect1[1])
	vcell, _ = CoordinatesToCellName(rect1[2], rect1[3])

	ws, err := f.workSheetReader(sheet)
	if err != nil {
		return err
	}
	ref := hcell + ":" + vcell
	if ws.MergeCells != nil {
		for i := 0; i < len(ws.MergeCells.Cells); i++ {
			cellData := ws.MergeCells.Cells[i]
			if cellData == nil {
				continue
			}
			cc := strings.Split(cellData.Ref, ":")
			if len(cc) != 2 {
				return fmt.Errorf("invalid area %q", cellData.Ref)
			}

			rect2, err := f.areaRefToCoordinates(cellData.Ref)
			if err != nil {
				return err
			}

			// Delete the merged cells of the overlapping area.
			if isOverlap(rect1, rect2) {
				ws.MergeCells.Cells = append(ws.MergeCells.Cells[:i], ws.MergeCells.Cells[i+1:]...)
				i--

				if rect1[0] > rect2[0] {
					rect1[0], rect2[0] = rect2[0], rect1[0]
				}

				if rect1[2] < rect2[2] {
					rect1[2], rect2[2] = rect2[2], rect1[2]
				}

				if rect1[1] > rect2[1] {
					rect1[1], rect2[1] = rect2[1], rect1[1]
				}

				if rect1[3] < rect2[3] {
					rect1[3], rect2[3] = rect2[3], rect1[3]
				}
				hcell, _ = CoordinatesToCellName(rect1[0], rect1[1])
				vcell, _ = CoordinatesToCellName(rect1[2], rect1[3])
				ref = hcell + ":" + vcell
			}
		}
		ws.MergeCells.Cells = append(ws.MergeCells.Cells, &xlsxMergeCell{Ref: ref})
	} else {
		ws.MergeCells = &xlsxMergeCells{Cells: []*xlsxMergeCell{{Ref: ref}}}
	}
	ws.MergeCells.Count = len(ws.MergeCells.Cells)
	return err
}

// UnmergeCell provides a function to unmerge a given coordinate area.
// For example unmerge area D3:E9 on Sheet1:
//
//    err := f.UnmergeCell("Sheet1", "D3", "E9")
//
// Attention: overlapped areas will also be unmerged.
func (f *File) UnmergeCell(sheet string, hcell, vcell string) error {
	ws, err := f.workSheetReader(sheet)
	if err != nil {
		return err
	}
	rect1, err := f.areaRefToCoordinates(hcell + ":" + vcell)
	if err != nil {
		return err
	}

	// Correct the coordinate area, such correct C1:B3 to B1:C3.
	_ = sortCoordinates(rect1)

	// return nil since no MergeCells in the sheet
	if ws.MergeCells == nil {
		return nil
	}

	i := 0
	for _, cellData := range ws.MergeCells.Cells {
		if cellData == nil {
			continue
		}
		cc := strings.Split(cellData.Ref, ":")
		if len(cc) != 2 {
			return fmt.Errorf("invalid area %q", cellData.Ref)
		}

		rect2, err := f.areaRefToCoordinates(cellData.Ref)
		if err != nil {
			return err
		}

		if isOverlap(rect1, rect2) {
			continue
		}
		ws.MergeCells.Cells[i] = cellData
		i++
	}
	ws.MergeCells.Cells = ws.MergeCells.Cells[:i]
	ws.MergeCells.Count = len(ws.MergeCells.Cells)
	return nil
}

// GetMergeCells provides a function to get all merged cells from a worksheet
// currently.
func (f *File) GetMergeCells(sheet string) ([]MergeCell, error) {
	var mergeCells []MergeCell
	ws, err := f.workSheetReader(sheet)
	if err != nil {
		return mergeCells, err
	}
	if ws.MergeCells != nil {
		mergeCells = make([]MergeCell, 0, len(ws.MergeCells.Cells))

		for i := range ws.MergeCells.Cells {
			ref := ws.MergeCells.Cells[i].Ref
			axis := strings.Split(ref, ":")[0]
			val, _ := f.GetCellValue(sheet, axis)
			mergeCells = append(mergeCells, []string{ref, val})
		}
	}

	return mergeCells, err
}

// MergeCell define a merged cell data.
// It consists of the following structure.
// example: []string{"D4:E10", "cell value"}
type MergeCell []string

// GetCellValue returns merged cell value.
func (m *MergeCell) GetCellValue() string {
	return (*m)[1]
}

// GetStartAxis returns the merge start axis.
// example: "C2"
func (m *MergeCell) GetStartAxis() string {
	axis := strings.Split((*m)[0], ":")
	return axis[0]
}

// GetEndAxis returns the merge end axis.
// example: "D4"
func (m *MergeCell) GetEndAxis() string {
	axis := strings.Split((*m)[0], ":")
	return axis[1]
}
