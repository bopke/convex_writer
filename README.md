# Convex writer

Draw simple convex shapes on your favorite picture

Send file via curl with desired color and vertices of quadrilateral figure like so:

```bash
curl -X POST -F file=@image.jpg \
-F "color=[64,32,128,0]" \
-F "vertices=[[400,400],[200,200],[400,250],[300,100]]" \
--output result.jpeg \
127.0.0.1:8000
```

This transforms such beautiful image:
![Beautiful screenshot of Haven and Hearth character](https://github.com/bopke/convex_writer/blob/master/image.jpg?raw=true "Beautiful screenshot of Haven and Hearth character")

Into such ugly mess:
![Convexelly vandalized mess](https://github.com/bopke/convex_writer/blob/master/result.jpeg?raw=true "Vandalized Haven and Hearth character")
