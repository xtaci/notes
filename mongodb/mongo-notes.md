mongos> var doc = db.BATTLES.find().sort({timestamp:-1}).limit(1)     
mongos> Object.bsonsize(doc.next())    
495
