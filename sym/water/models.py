from django.db import models
 
class Manage(models.Model):
    service     = models.CharField('应用',  max_length=20)#服务
    name        = models.CharField('名称',  max_length=20)#文件或者啥
    token       = models.CharField('token', max_length=20)#访问应用的token
    message     = models.TextField('备注') #备注
    status      = models.CharField('状态', max_length=20) 
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 

    def __str__(self):
        return self.name
    
    class Meta:
        verbose_name = '水印管理'
        verbose_name_plural = '水印管理'


class Log(models.Model):
    service     = models.CharField('服务',  max_length=20)#服务
    upload      = models.CharField('上传文件名',  max_length=20)
    status      = models.CharField('状态', max_length=20) #密钥生成、更新、注销、归档、删除
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 

    class Meta:
        verbose_name = '水印日志'
        verbose_name_plural = '水印日志'


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






