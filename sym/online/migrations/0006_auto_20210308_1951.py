# Generated by Django 3.1.7 on 2021-03-08 11:51

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('online', '0005_auto_20210308_1919'),
    ]

    operations = [
        migrations.AlterField(
            model_name='online',
            name='disk',
            field=models.CharField(default='100%', editable=False, max_length=200, verbose_name='总计硬盘空间'),
        ),
    ]
