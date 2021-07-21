# nubeio-rubix-app-lora-go

# Tables
## network/device/point
- each network/device/point can have one or more tags
```
-- network -> tags
--- device -> tags 
---- point  -> tags -> thingRef
```
## pre-made device
- a pre-made device is a pre-made device with one or more points (for example a droplet)

## tags
A tag can be added to any net/dev/point

## thingRef
A thingRef is an association to a real world object. Like a 

# Goal's

## Have a way for front end to know what the can and can't CRUD
for example
```
        "point": [
            {
                "uuid": "ae243e782fe7411d",
                "name": "aidan is cool1111",
                "dataType": "readHoldingRegister",
            }, "ops"{ //opertaions allowed 
                "uuid": {"write":false, "required":true},
                "name": {"write":true, "required":true, "dataType":"str"}
                "dataType": {"write":true, "required":true, "dataType":{"readHoldingRegister":readHoldingRegister, "writeHoldingRegister":writeHoldingRegister}}
            }

```

# maybe grpc
