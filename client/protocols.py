#!/usr/bin/python
# -*- coding: utf-8 -*-
#import json
import base64,urllib,httplib,json,os
import pexpect 
import getpass 
import commands
import time
import sys
import paramiko
from urlparse import urlparse
sys.path.append('/home/yangxin/servicePlatform/servicePlatform')
#from jsonDecoders import JsonDecoder
# from models import API
# from application.models import application


#def singleton(cls, *args, **kw):
#        instances = {}
#        def _singleton():
#                if cls not in instances:
#                        instances[cls] = cls(*args, **kw)
#                return instances[cls]
#        return _singleton
#
#@singleton
#class OpenStack(API):
#        a = 1
#        def __init__(self, x=0):
#                self.x = x
#
#        def do(self,_name,_token,_host,_data='',_tenantId=''):
#                api = API.objects.get(name=_name)
#                if(''!=_token):
#                        api.head = api.head% _token
#                if(''!=_tenantId):
#                        api.url = api.url % _tenantId
#                if(0!=len(_data)):
#                        api.data = api.data % _data
#                api.host = _host
#
#                return send(api)
#
#
#def send(api):
#        conn = httplib.HTTPConnection(api.host,api.port)
#        conn.request(api.method,api.url,api.data,eval(api.head))
#        response  = conn.getresponse()
#        #
#        #response = trim(response)
#        status = response.status
#        reason = response.reason
#        #if(status != "200")
#                #return status+" "+reason
#        heads = response.getheaders()
#        body = response.read()
#        conn.close()
#        body = JsonDecoder().getToken(body)
#        return body;
#        #return response;

class GetPara():
	def __init__(self, x=0):  
		self.x = x  
	def do(self,name,imgId,disk,ip,username,password):
		dd = getToken();
		instanceid = creatInstance(dd,name,imgId,disk,ip,username,password);		
		return instanceid

	
def getToken():
	url1="10.109.252.247:5000"
	params1 = '{"auth": {"tenantName": "admin", "passwordCredentials":{"username": "admin", "password": "ADMIN_PASS"}}}'
	headers1 = {"Content-Type": 'application/json'}
	conn1 = httplib.HTTPConnection(url1)
	conn1.request("POST","/v2.0/tokens",params1,headers1)
	response1 = conn1.getresponse()
	data1 = response1.read()
	dd1 = json.loads(data1)
        print "token",dd1
	return dd1

def creatInstance(dd,name,imgId,disk,ip,username,password):
	url4 = "10.109.252.247:8774"
	#flavorRef:1   ===> m1.tiny
	imageRef = imgId
	networkid = '3d68e0e6-6773-455e-9850-88d3ace41490'
	#networkid = '3c6c13fb-75ec-4065-b557-7a9c64ce663d'
	if('1GB' == disk):
		flavorRef = '1'
	elif('20GB' == disk):
		flavorRef = '2'
	elif('40GB' == disk):
		flavorRef = '3'
	else:
		flavorRef = '3'

#	puppet_scripts = '/usr/bin/puppet agent --server puppet.domain.com --test'
	#param_tpl = "#!/bin/sh\npasswd %s<<EOF\n%s\n%s\nEOF"  
	param_tpl = """#!/bin/sh
passwd %s<<EOF
%s
%s
EOF
sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/g' /etc/ssh/sshd_config
service ssh restart
""" 
	
	puppet_scripts = param_tpl % (username,password,password)
	userdata=base64.b64encode(puppet_scripts)
	param_template = '{"server": {"key_name":"mykey","name":"%s", "imageRef":"%s", "flavorRef":"%s", "networks":[{"uuid":"%s"}], "user_data":"%s", "max_count":1, "min_count":1}}'
#	params4 = '{"server": {"name":"'+app_name+'", "imageRef": "'+imageRef+'","flavorRef": "'+flavorRef+'", "max_count": 1, "min_count": 1}}'
	params4 = param_template  % (name, imageRef, flavorRef, networkid, userdata)
	headers4 = { "X-Auth-Token":dd['access']['token']['id'], "Content-type":"application/json" }
	conn4 = httplib.HTTPConnection(url4)
	#conn4.request("POST", "/v2.1/%s/servers" % urlparse(dd['access']['serviceCatalog'][0]['endpoints'][0]['publicURL'])[2], params4, headers4)
	conn4.request("POST", "/v2.1/servers", params4, headers4)
	# print dd['access']['token']['id']
	# print urlparse(dd['access']['serviceCatalog'][1]['endpoints'][0]['publicURL'])[2]

	response4 = conn4.getresponse()
	data4 = response4.read()
	dd4 = json.loads(data4)
	print json.dumps(dd4, sort_keys=True, indent=4)
	if dd4.has_key("server"):
                print "has server"
		ip = AllocateIp().do(dd4["server"]["id"],ip)
                print "ip",ip
		if ip =="":
			return None
		return dd4["server"]["id"]
	return None

class AllocateIp():
	def do(self,instanceid,ip): 
		isSuccess = notCreated(instanceid)
		if not isSuccess:
			return ""
                dd = getToken();
                url = "10.109.252.247:8774"
                params = '{"addFloatingIp" : {"address": "%s"}}'%ip
                headers = { "X-Auth-Token":dd['access']['token']['id'], "Content-type":"application/json" }
                conn= httplib.HTTPConnection(url)
                conn.request("POST", "/v2.1/servers/%s/action" % instanceid, params, headers)
		#cmd = "nova add-floating-ip "+instanceid+" "+ip
		#print cmd
		#child = pexpect.spawn(cmd)
		#data = child.read()
		#print data
		#child.close()
		
		return ip

def notCreated(instanceid):
	count = 0
	while True:
		dd = getToken();
		url5 = "10.109.252.247:8774"
		params5 = urllib.urlencode({})
		headers5 = { "X-Auth-Token":dd['access']['token']['id'], "Content-type":"application/json" }
		conn5= httplib.HTTPConnection(url5)
		conn5.request("GET", "/v2.1/%s/servers/%s" % (urlparse(dd['access']['serviceCatalog'][0]['endpoints'][0]['publicURL'])[2],instanceid), params5, headers5)
		response5 = conn5.getresponse()
		data5= response5.read()
		dd5= json.loads(data5)
		print json.dumps(dd5, sort_keys=True, indent=4)
		if dd5["server"]["status"] == "ACTIVE":
			break
		# print "be there"
		time.sleep(1)
		count = count + 1
		if count > 150 :
			return False
	return True

class deleteInstance():
        def do(self,instanceid):
                dd = getToken();
                url = "10.109.252.247:8774"
                params = urllib.urlencode({})
                headers = { "X-Auth-Token":dd['access']['token']['id'], "Content-type":"application/json" }
                conn = httplib.HTTPConnection(url)
                conn.request("DELETE", "%s/servers/%s" % (urlparse(dd['access']['serviceCatalog'][0]['endpoints'][0]['publicURL'])[2],instanceid), params, headers)
                response = conn.getresponse()


class doScp():
        def do(self,user,password, host, path, files):
                fNames = files
                name = files.split("/")[-1]
                child = pexpect.spawn('scp %s %s@%s:%s' % (fNames, user, host,path))
                i = child.expect(['password:', r"yes/no"], timeout=30)
                if i == 0:
                        child.sendline(password)
                elif i == 1:
                        child.sendline("yes")
                        child.expect("password:", timeout=30)
                        child.sendline(password)
                data = child.read()
                child.close()
                return (data.find(name)  != -1)

# add by @lvp
def delApp(ip, username, password, pkgname):
        try:
            ssh = paramiko.SSHClient()
            ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            ssh.connect(ip, 22, username, password, timeout=30)
            for m in pkgname:
                cmd = 'rm /usr/local/tomcat/webapps/' + m
                print cmd
                stdin, stdout, stderr = ssh.exec_command(cmd)
            print '%s\tOK\n' % ip
            ssh.close()
        except:
                print '%s\tError\n' % ip


def getImageID():
	token = getToken()
	url = '10.109.252.247:9292'
	params = urllib.urlencode({})
#	headers = {"Content-Type": "application/json", "X-Auth-Token":token['access']['token']['id'] }
	headers = {"X-Auth-Project-Id": "admin", "Content-Type": "application/json", "Accept": "application/json", "X-Auth-Token":token['access']['token']['id'] }
	conn = httplib.HTTPConnection(url)
	conn.request("GET", "/v1/images", params, headers)
	response = conn.getresponse()
	data = response.read()
	imageInfo = json.loads(data)
#	print json.dumps(imageInfo, sort_keys=True, indent=4)

	imageID = dict()
	flag = -1
	for image in imageInfo['images']:
		for key in image:
			if key == 'name':
				name = image[key]
				flag += 1
			if key == 'id':
				uid = image[key]
				flag += 1
			if flag > 0:
				imageID[name] = uid
   		
	return imageID

def IsActive(myname):
	cnt = 0
	while True:
		token = getToken()
		url = '10.109.252.247:9292'
		params = urllib.urlencode({})
	#	params = '{"detail":{"name": "%s"}}' % myname
	#	headers = {"Content-Type": "application/json", "X-Auth-Token":token['access']['token']['id'] }
		headers = {"X-Auth-Project-Id": "admin", "Content-Type": "application/json", "Accept": "application/json", "X-Auth-Token":token['access']['token']['id'] }
		conn = httplib.HTTPConnection(url)
		conn.request("GET", "/v1/images/detail", params, headers)
		response = conn.getresponse()
		data = response.read()
#	imageInfo = json.loads(data)
#	print json.dumps(imageInfo, sort_keys=True, indent=4)
		flag = False
		detail = dict()
	
		for image in imageInfo['images']:
			if flag:
				break
			for key in image:
				if key == 'name':
					if image[key] == myname:
						detail = image
						flag = True
						break
		if(detail['status'] == 'active'):
#	print json.dumps(detail, sort_keys=True, indent=4)
			return True
		if(detail == {}):
			return False
		time.sleep(2)
		cnt += 1
		if cnt > 300:
			return False
	

	return False
   		

def createImage(instance_id, myname):
    token = getToken()
    url = '10.109.252.253:8774'
    params = '{"createImage": {"name": "%s", "metadata": {}}}' % myname
    headers = {"X-Auth-Project-Id": "admin", "Content-Type": "application/json", "Accept": "application/json", "X-Auth-Token":token['access']['token']['id'] }
    conn = httplib.HTTPConnection(url)
    conn.request("POST", "%s/servers/%s/action" % (urlparse(token['access']['serviceCatalog'][0]['endpoints'][0]['publicURL'])[2], instance_id), params, headers)
	
    time.sleep(1)
    imageID = getImageID()
    return imageID[myname]

if __name__ == '__main__':
	token = getToken()
	name = "test"
	imgId = "5403048b-8162-447c-8304-b671ede4031b"
	disk = "40GB"
	ip = '10.109.252.195'
	username = 'ubuntu'
	password = 'ubuntu'
	x = creatInstance(token, name, imgId, disk, ip, username, password)
	print x
	headers = { "X-Auth-Token":token['access']['token']['id'], "Content-type":"application/json" }
        params=urllib.urlencode({})
	conn4 = httplib.HTTPConnection("10.109.252.247:8774")
        conn4.request("GET", "/v2.1/servers", params, headers)
        response = conn4.getresponse()
        data = response.read()
	print data
	#name = 'test'
        #IsActive(name)
	
#	imageID = getImageID()
#	print json.dumps(imageID, sort_keys=True, indent=4), len(imageID)
#	instance_id = '7319871e-0ff2-4f3d-9dfd-74cae32599f6'
#	myname = "lvpsnapshot"
#	print urlparse(token['access']['serviceCatalog'][1]['endpoints'][0]['publicURL'])[2]
#	result = createImage(instance_id, myname)
#	print result
