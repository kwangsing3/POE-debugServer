# POE-debugServer
It's a data Server for Developing a data stats server which service for Path of exile.

It's made for simulating the way how Path of exile offical server gives data.
For debugging our applications or websites, either trade tool.
Only set with internally default character data or trade data, without update data from Offical server.

(Which mean it's a localhost server used for confirming your data struct is right or not, also helping you to debug your application to make sure your handle is right if the server gave the unexpect struct or shut down for maintenance.
Also have a board to contral the debug server to different status and ensure your application errors handle is fine.)



Expect feature:
- maintenance control                            (to test when the server shut down accidentally)
- maintenance control with timer
- API simulate with contralable results          (including trade and character data.)
- Gives illegal data struct randomly.            (including gives depercated data, maybe?)
- .... etc
