import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import Manage, Log
#from flyadmin.widget.forms import SelectBoxWidget, TimelineWidget, EditorWidget, DateTimeWidget, UploadImagesWidget, InputNumberWidget, UploadFileWidget, StepsWidget, StepsNormalWidget
 

class ManageAdminForm(forms.ModelForm):
  class Meta:
      model = Manage
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(ManageAdminForm, self).__init__(*args, **kwargs)
  
class ManageAdmin(admin.ModelAdmin):
    form = ManageAdminForm 
    list_display = ('service', 'name', 'status', 'update_time')
    search_fields = ('name','service',)

admin.site.register(Manage, ManageAdmin)



class LogAdminForm(forms.ModelForm):
  class Meta:
      model = Log
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(LogAdminForm, self).__init__(*args, **kwargs)
  
class LogAdmin(admin.ModelAdmin):
    form = LogAdminForm 
    list_display = ('service', 'status','update_time')
    search_fields = ('name','service',)

admin.site.register(Log, LogAdmin)


