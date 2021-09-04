from django.shortcuts import render
from django.apps import apps


#所有实时数据概览、数据详情、数据分布、
#敏感数据概况、敏感数据分布、任务数等各类数据统计、
#数据源分级分类统计，分析维度包含：数据库总表、发现数据表、敏感数据表、数据分级、数据分类

 
def admin_index(request):
    #将数据写入到里面来到首页数据里(不得修改)
    html = ['<el-row>']
    for app in apps.get_app_configs():
        for model in app.get_models():
            if hasattr(model, 'show_plots'):
                show_plots = model.show_plots()
                if show_plots:
                    for plot in show_plots:
                        html.append('<div class="el-col el-col-{}">{}</div>'.format(plot.get('size'), plot.get('plot')))
    html.append('</el-row>')
    xplots = ''.join(html)   
    return render(request, "admin/xplots.html", locals()) 
 