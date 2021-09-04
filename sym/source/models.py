from django.db import models


class Encoder(models.Model):
    service = models.CharField('服务', max_length=20)#服务
    upload  = models.FileField('上传文件')
    message = models.CharField('水印数据', max_length=200)
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 

    class Meta:
        verbose_name = '新增水印'
        verbose_name_plural = '新增水印'

class Decoder(models.Model):
    service = models.CharField('服务',  max_length=20)#服务
    upload  = models.FileField('上传文件')
    message = models.CharField('水印数据', max_length=200)
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 
    
    class Meta:
        verbose_name = '水印溯源'
        verbose_name_plural = '水印溯源'