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