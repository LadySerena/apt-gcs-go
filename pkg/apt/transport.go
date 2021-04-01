package apt

//basic flow
//init gcs client on start
//send out 100 capabilities message to apt on startup
//receive 600 from apt
//send out 200 start message
//get the thing from gcs
//send 201 if happy or 400 if sad
//there's something weird with the encoding per https://github.com/dhaivat/apt-gcs/issues/1