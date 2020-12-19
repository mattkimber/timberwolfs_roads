#!/bin/bash

# Necessary directories
mkdir -p processed/slopes
mkdir -p processed/catenary_slopes
mkdir -p processed/depots


# Processed objects
echo "Compositing"
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/straight positor/slopes.json
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/bridge positor/slopes.json

../cargopositor/cargopositor.exe -o processed/catenary_slopes -v voxels/catenary/straight positor/slopes.json
../cargopositor/cargopositor.exe -o processed/catenary_slopes -v voxels/catenary/bridge positor/slopes.json

../cargopositor/cargopositor.exe -o processed/depots -v voxels/depot/regular positor/depots.json
../cargopositor/cargopositor.exe -o processed/depots -v voxels/depot/electric positor/depots_elec.json

../cargopositor/cargopositor.exe -o processed/slopes -v voxels/rhs/straight positor/slopes.json
../cargopositor/cargopositor.exe -o processed/slopes -v voxels/rhs/bridge positor/slopes.json


# Render
echo "Rendering roads/rails"
../gorender/renderobject.exe -m files/manifest_1x.json -8 -s 1,2 -u -r -progress voxels/crossroad/*.vox

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r -progress voxels/junctions/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/straight/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/no_inclines/*.vox

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r -progress voxels/misc/*.vox

## RHS

../gorender/renderobject.exe -m files/manifest_1x.json -8 -s 1,2 -u -r -progress voxels/rhs/crossroad/*.vox

../gorender/renderobject.exe -m files/manifest_4x.json -8 -s 1,2 -u -r -progress voxels/rhs/junctions/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/rhs/straight/*.vox

../gorender/renderobject.exe -m files/manifest_2x.json -8 -s 1,2 -u -r -progress voxels/rhs/no_inclines/*.vox


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

echo
echo "Rendering depots"


../gorender/renderobject.exe -m files/manifest_depot_4x.json -8 -s 1,2 -u -r -progress processed/depots/*.vox



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
../splatter/splatter.exe -i 2x -o sheets_2x -d files/spritesheet_catenary.json -m 8 

# Depot spritesheets
../splatter/splatter.exe -i 1x -o sheets_1x -d files/spritesheet_depot.json -m 4 
../splatter/splatter.exe -i 2x -o sheets_2x -d files/spritesheet_depot.json -m 8 


# Run Roadie
echo "Creating NML"
../roadie/roadie.exe set.json

# NML compilation
echo "Compiling NML"
../nml/nmlc.exe -c timberwolfs_roads.nml

# Build TAR
echo "Building TAR"
mkdir -p timberwolfs_roads
mv *.grf timberwolfs_roads
cp grf_readme/* timberwolfs_roads
tar -c timberwolfs_roads > timberwolfs_roads.tar