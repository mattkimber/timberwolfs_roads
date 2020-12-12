#!/bin/bash

# Necessary directories
mkdir -p processed/slopes

# Processed objects
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/straight positor/slopes.json

# Render
../gorender/renderobject.exe -m files/manifest_1x.json -8 -s 1,2 -u -r voxels/crossroad/*.vox

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r voxels/junctions/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r voxels/straight/*.vox
../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r processed/slopes/*.vox

# Assemble spritesheets
mkdir -p sheets_1x
mkdir -p sheets_2x

# Spritesheets with mask
../splatter/splatter.exe -i 1x -o sheets_1x -d files/spritesheet.json -m 4 -k files/mask_1x.png
../splatter/splatter.exe -i 2x -o sheets_2x -d files/spritesheet.json -m 8 -k files/mask_2x.png

# Run Roadie
../roadie/roadie.exe set.json

# NML compilation
../nml/nmlc.exe timberwolfs_roads.nml