#!/bin/bash

# Necessary directories
mkdir -p processed/slopes

# Processed objects
echo "Compositing"
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/straight positor/slopes.json
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/bridge positor/slopes.json

# Render
echo "Rendering"
../gorender/renderobject.exe -m files/manifest_1x.json -8 -s 1,2 -u -r -progress voxels/crossroad/*.vox

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r -progress voxels/junctions/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/straight/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/no_inclines/*.vox

# Don't render the bridge/ directory as it contains only ramps

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r -progress processed/slopes/*.vox


# Assemble spritesheets
echo "Assembling Spritesheets"
mkdir -p sheets_1x
mkdir -p sheets_2x

# Spritesheets with mask
../splatter/splatter.exe -i 1x -o sheets_1x -d files/spritesheet.json -m 4 -k files/mask_1x.png
../splatter/splatter.exe -i 2x -o sheets_2x -d files/spritesheet.json -m 8 -k files/mask_2x.png

# Run Roadie
echo "Creating NML"
../roadie/roadie.exe set.json

# NML compilation
echo "Compiling NML"
../nml/nmlc.exe timberwolfs_roads.nml