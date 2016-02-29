#!/usr/bin/env python
# -*- coding: utf-8 -*-

from django.conf.urls import url, include
from rest_framework.routers import DefaultRouter
from views import TodoItemViewSet

__author__ = "wangyixiang"

dr = DefaultRouter()
dr.register("item", TodoItemViewSet)

urlpatterns = [
    url(r'^', include(dr.urls)),
]
