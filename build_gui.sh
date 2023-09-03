#!/bin/bash
# GUI
echo
echo "Creating GUI"


# GUI
mkdir -p gui_rendered
mkdir -p gui_rendered/1x
mkdir -p gui_rendered/2x
mkdir -p gui_output
mkdir -p gui_output/1x
mkdir -p gui_output/2x

# Depot GUI
../gorender/renderobject.exe -o gui_rendered/1x/depot -8 -s 1 -m files/manifest_gui.json voxels/depot/empty/foster.vox 
../gorender/renderobject.exe -o gui_rendered/2x/depot -8 -s 2 -m files/manifest_gui.json voxels/depot/empty/foster.vox 

# Roads
for i in voxels/crossroad/*.vox; do 
    j=$(echo $i | cut -f3 -d/)z
    ../gorender/renderobject.exe -o gui_rendered/1x/$j -8 -s 1 -m files/manifest_gui_crossroad.json $i
    ../gorender/renderobject.exe -o gui_rendered/2x/$j -8 -s 2 -m files/manifest_gui_crossroad.json $i
done  

for i in voxels/straight/*.vox; do 
    j=$(echo $i | cut -f3 -d/)
    ../gorender/renderobject.exe -o gui_rendered/1x/$j -8 -s 1 -m files/manifest_gui_road.json $i
    ../gorender/renderobject.exe -o gui_rendered/2x/$j -8 -s 2 -m files/manifest_gui_road.json $i
    k=$(echo $j | cut -f1 -d.)
    ./roadgoo/roadgoo.exe $k
done  