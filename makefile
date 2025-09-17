.PHONY: clean export assets build
clean:
	rm -rf dist
export:
	go run .
assets:
	# copy everything in public/ to dist/
	# if your CSS is in public/style.css and your HTML links to /assets/style.css, place it accordingly:
	mkdir -p dist/assets
	# If your assets are in public/assets/* already:
	# cp -R public/assets dist/
	# Or, if you keep a flat public/style.css:
	cp public/style.css dist/assets/style.css
	# Copy any other static files (images, docs, etc.)
	[ -d public ] && cp -R public/* dist/ || true
build: clean export assets
	@echo "Built static site in ./dist"
