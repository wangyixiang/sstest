#!/usr/bin/env python
# -*- coding: utf-8 -*-

from rest_framework.serializers import ModelSerializer
from models import TodoItem

__author__ = "wangyixiang"


class TodoItemSerializer(ModelSerializer):
    class Meta:
        model = TodoItem
