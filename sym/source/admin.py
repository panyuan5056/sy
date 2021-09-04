import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import  Encoder, Decoder
# Register your models here.


class EncoderAdminForm(forms.ModelForm):
  class Meta:
      model = Encoder
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(EncoderAdminForm, self).__init__(*args, **kwargs)
  
class EncoderAdmin(admin.ModelAdmin):
    form = EncoderAdminForm 
    list_display = ('service', 'upload','update_time')
    search_fields = ('service','upload',)

admin.site.register(Encoder, EncoderAdmin)




class DecoderAdminForm(forms.ModelForm):
  class Meta:
      model = Decoder
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(DecoderAdminForm, self).__init__(*args, **kwargs)
  
class DecoderAdmin(admin.ModelAdmin):
    form = DecoderAdminForm 
    list_display = ('service', 'upload','update_time')
    search_fields = ('service','upload',)

admin.site.register(Decoder, DecoderAdmin)