from rest_framework.viewsets import ModelViewSet
from rest_framework.authentication import SessionAuthentication
from rest_framework.permissions import AllowAny
from serializers import TodoItemSerializer
from models import TodoItem


# Create your views here.

class TodoItemViewSet(ModelViewSet):
    serializer_class = TodoItemSerializer
    queryset = TodoItem.objects.all()
    authentication_classes = (SessionAuthentication,)
    permission_classes = (AllowAny,)
