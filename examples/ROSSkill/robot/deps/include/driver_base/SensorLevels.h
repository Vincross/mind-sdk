#ifndef _ROS_driver_base_SensorLevels_h
#define _ROS_driver_base_SensorLevels_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "ros/msg.h"

namespace driver_base
{

  class SensorLevels : public ros::Msg
  {
    public:
      enum { RECONFIGURE_CLOSE =  3   };
      enum { RECONFIGURE_STOP =  1   };
      enum { RECONFIGURE_RUNNING =  0  };

    SensorLevels()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
     return offset;
    }

    const char * getType(){ return "driver_base/SensorLevels"; };
    const char * getMD5(){ return "6322637bee96d5489db6e2127c47602c"; };

  };

}
#endif