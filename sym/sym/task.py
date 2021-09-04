
import json
import requests
import psutil
from task.models import TaskDetail, Task
from online.models import Online
from natural.models import Natural
from datetime import datetime as dt
 
def my_scheduled_job():

    #更新cpu
    cpu = psutil.cpu_percent()
    Natural(**{'name':'cpu', 'ava':cpu, 'total':'1'}).save()

    #更新online cpu
    for online in Online.objects.order_by("cpu"):
        try:
            body = requests.post(q.server + '/online', headers={"token":online.token})
            if body.status_code == 200:
                data = body.json()
                online.status = 2
                online.cpu    = '{}%'.format(int(data['result']['cpu']))
                online.memery = '{}%'.format(int(data['result']['memery']))
                online.disk   = '{}T'.format(round(data['result']['disk']['total']/(1024*1024*1024*1024),2))
                online.disk2  = '{}T'.format(round(data['result']['disk']['free']/(1024*1024*1024*1024),2))
            else:
                online.status = 3
        except Exception as e:
            print(e)
            online.status = 3
            online.update_time = dt.now()
        online.save()

    #获取发现内容
    tasks = Task.objects.filter(status = 2).all()
    for task in tasks:
        task_detail = task_detail.objects.filter(task= task).first()
        if task_detail:
            version = task_detail + 1 
        else:
            version = 1
        for online in Online.objects.order_by("cpu"):
            try:
                payload = {'id':task.tid}
                body = requests.post(online.server + '/api/v1/get', headers={"token":online.token}, data=json.dumps(payload))
                if body.status_code == 200:
                    result = body.json()
                    if result.get('status') == 200:
                        reports = json.loads(result['result']['Report'])
                        for tablename, report in reports.items():
                            for col in report:
                                col['version'] = version
                                col['task'] = task
                                col['table_name'] = tablename
                                TaskDetail(**col).save()
                        task.status = 3
                        task.version = version
                        task.save()
                break
            except Exception as e:
                print(e)
                 