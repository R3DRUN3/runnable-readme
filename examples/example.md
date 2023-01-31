# THIS IS AN EXAMPLE MARKDOWN FILE

This is a console identifier:  
```console
echo Hello World!
```

<br/>

This is another console identifier:  
```console
export IP=$(curl ifconfig.me) \
&& echo "My external IP is: $IP" \
&& echo Sleeping for 3 seconds \
&& sleep 3 \
&& echo Goodbye!
```
<br/>

And another one!  
```console
cd /tmp \
&& echo "where am I?" \
&& pwd \
&& echo  "let's create folders in loop" \
&& for i in {1..10}; do mkdir temp_repo_$i; done \
&& echo "Sleeping for 2 seconds" \
&& sleep 2 \
&& echo "removing the folders" \
&& rm -r temp_repo_*
```
<br/>




