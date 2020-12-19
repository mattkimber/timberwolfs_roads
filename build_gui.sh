#!/bin/bash
# GUI
echo
echo "Creating GUI"


# GUI
mkdir -p gui_rendered

# Depot GUI
../gorender/renderobject.exe -o gui_rendered/depot -8 -s 1 -m files/manifest_gui.json voxels/depot/gui/foster.vox 

# Roads
for i in voxels/crossroad/*.vox; do 
    j=$(echo $i | cut -f3 -d/)
    ../gorender/renderobject.exe -o gui_rendered/$j -8 -s 1 -m files/manifest_gui_crossroad.json $i
done  

for i in voxels/straight/*.vox; do 
    j=$(echo $i | cut -f3 -d/)
    ../gorender/renderobject.exe -o gui_rendered/$j -8 -s 1 -m files/manifest_gui_road.json $i
    k=$(echo $j | cut -f1 -d.)
    ./roadgoo/roadgoo.exe $k
done  