ROS Skill
====

ROS Skill implements a *skill* that shows how to publish images to an ROS topic via rosserial.

Prerequisites
----

1. Install [ROS system](http://www.ros.org/install) and [rosserial](http://wiki.ros.org/rosserial_embeddedlinux/GenericInstall) on your PC.

Building and running the example
----

Launch [roscore](http://wiki.ros.org/roscore) on your PC:

```
$ roscore
```

To have your PC receive ROS messages from your ROBOT, execute the following command in a terminal on your PC:
```
$ rosrun rosserial_python serial_node.py tcp
```

Modify `rosMasterIP` inside of `robot/src/rosdemoskill.go` to point to your PC's IP.

You should now be ready to build, pack and run the *Skill*
```
$ mind build && mind pack && mind run
```

Once your *Skill* is running you can see the images being streamed to your PC by executing:
```
$ rosrun image_view image_view compressed   # for ros indigo
$ rosrun rqt_image_view rqt_image_view      # for ros kinetic
```

*You might need to install [compressed_image_transport](http://wiki.ros.org/compressed_image_transport).*

Notes
----
This example *Skill* has only been tested on ROS indigo and kinetic.

If you need to add other message types than those already provided you can follow the official guide on how to [generate rosserial_embedded linux libraries](http://wiki.ros.org/rosserial_embeddedlinux/GenericInstall).
Example:
```
cd <some_directory>
rm -rf ros_lib examples
rosrun rosserial_embeddedlinux make_libraries.py .
```

Finally, put the `ros_lib` to `robot/deps/include` in skill folder.

Limitations
----
[rosserial](http://wiki.ros.org/rosserial) uses 2 bytes to describe the length of the message, thus, the length of messages can not exceed 32767 bytes. 
However, you can try to write the sender and receiver part by yourself to support larger messages.
