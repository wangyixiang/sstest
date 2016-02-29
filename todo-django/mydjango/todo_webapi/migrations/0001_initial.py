# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import models, migrations


class Migration(migrations.Migration):

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='TodoItem',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('content', models.CharField(max_length=1024)),
                ('created_time', models.DateTimeField(auto_now_add=True)),
                ('finished_time', models.DateTimeField(auto_now=True)),
                ('finished', models.BooleanField()),
            ],
        ),
    ]
