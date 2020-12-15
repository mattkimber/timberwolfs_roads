#!/bin/bash

# Necessary directories
mkdir -p processed/slopes
mkdir -p processed/catenary_slopes


# Processed objects
echo "Compositing"
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/straight positor/slopes.json
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/bridge positor/slopes.json

../cargopositor/cargopositor.exe -o processed/catenary_slopes -v voxels/catenary/straight positor/slopes.json
../cargopositor/cargopositor.exe -o processed/catenary_slopes -v voxels/catenary/bridge positor/slopes.json

# Render
echo "Rendering roads/rails"
../gorender/renderobject.exe -m files/manifest_1x.json -8 -s 1,2 -u -r -progress voxels/crossroad/*.vox

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r -progress voxels/junctions/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/straight/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/no_inclines/*.vox

# Don't render the bridge/ directory as it contains only ramps

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r -progress processed/slopes/*.vox

echo
echo "Rendering catenary"
../gorender/renderobject.exe -m files/manifest_catenary_1x.json -8 -s 1,2 -u -r -progress voxels/catenary/crossroad/*.vox

../gorender/renderobject.exe -m files/manifest_catenary_4x.json -8 -s 1,2 -u -r -progress voxels/catenary/junctions/*.vox

../gorender/renderobject.exe -m files/manifest_catenary_2x.json -8 -s 1,2 -u -r -progress voxels/catenary/straight/*.vox

../gorender/renderobject.exe -m files/manifest_catenary_2x.json -8 -s 1,2 -u -r -progress voxels/catenary/bridge_no_incline/*.vox

# Don't render the bridge/ directory as it contains only ramps

../gorender/renderobject.exe -m files/manifest_catenary_4x.json -8 -s 1,2 -u -r -progress processed/catenary_slopes/*.vox


# Assemble spritesheets
echo
echo "Assembling Spritesheets"
mkdir -p sheets_1x
mkdir -p sheets_2x

# Spritesheets with mask
../splatter/splatter.exe -i 1x -o sheets_1x -d files/spritesheet.json -m 4 -k files/mask_1x.png
../splatter/splatter.exe -i 2x -o sheets_2x -d files/spritesheet.json -m 8 -k files/mask_2x.png

# Catenary spritesheets
../splatter/splatter.exe -i 1x -o sheets_1x -d files/spritesheet_catenary.json -m 4 
#-k files/catenary_mask_1x.png
../splatter/splatter.exe -i 2x -o sheets_2x -d files/spritesheet_catenary.json -m 8 
#-k files/catenary_mask_2x.png

# Run Roadie
echo "Creating NML"
../roadie/roadie.exe set.json

# NML compilation
echo "Compiling NML"
../nml/nmlc.exe -c timberwolfs_roads.nml