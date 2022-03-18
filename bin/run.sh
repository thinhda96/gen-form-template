#!/bin/bash

echo Start zip settings
cd ~/.gen_form_template/settings 
zip -r settings.zip * 
cd - 
cp ~/.gen_form_template/settings/settings.zip .
echo Zip ok