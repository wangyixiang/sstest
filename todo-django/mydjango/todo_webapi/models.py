from django.db import models


# Create your models here.
class TodoItem(models.Model):
    content = models.CharField(max_length=1024)
    created_time = models.DateTimeField(auto_now_add=True)
    finished_time = models.DateTimeField(auto_now=True)
    finished = models.BooleanField()

    def __unicode__(self):
        return "created_time: " + self.created_time + "\n" + \
               "finised: " + self.finished + "\n" + \
               "finished_time: " + self.finished_time + "\n" + \
               "content: " + self.content


