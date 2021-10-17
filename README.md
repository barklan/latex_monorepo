# LaTeX monorepo

https://wiki.archlinux.org/title/TeX_Live#Making_fonts_available_to_Fontconfig

ln -s /usr/share/texmf-dist/fonts/opentype/public/stix2-otf/STIXTwoMath-Regular.otf ~/.local/share/fonts/OTF/

fc-cache ~/.local/share/fonts
mkfontscale ~/.local/share/fonts/OTF
mkfontdir ~/.local/share/fonts/OTF
