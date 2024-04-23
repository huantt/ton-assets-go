# Get the current version from the latest tag (default format vX.Y.Z)
current_version=$(git describe --abbrev=0 --tags 2>/dev/null)
# If no tags exist, start with v0.0.1
if [ -z "$current_version" ]; then
  current_version="v0.0.1"
fi

# Split the version into major, minor, and patch components
major="${current_version%%.*}"
major=${major:1}  # Remove leading "v"
minor="${current_version#*.}"
minor="${minor%%.*}"
patch="${current_version#*.*}"
patch="${patch##*.}"

# Get the increment type (major, minor, or patch) from the first argument (optional)
increment_type="${1:-patch}"

# Increment the version based on the increment type
case "$increment_type" in
  major)
    ((major++))
    minor="0"
    patch="0"
    ;;
  minor)
    ((minor++))
    patch="0"
    ;;
  patch)
    ((patch++))
    ;;
  *)
    echo "Invalid increment type. Valid options: major, minor, patch"
    return 1
    ;;
esac

# Build the new version string
new_version="v${major}.${minor}.${patch}"
echo "$new_version"