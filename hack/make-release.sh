

set -euo pipefail
SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
[[ -n "${DEBUG:-}" ]] && set -x

log() { echo "$1" >&2; }
fail() { log "$1"; exit 1; }

TAG="${TAG:?TAG env variable must be specified}"
REPO_PREFIX="${REPO_PREFIX:?REPO_PREFIX env variable must be specified e.g. gcr.io\/google-samples\/microservices-demo}"

if [[ "$TAG" != v* ]]; then
    fail "\$TAG must start with 'v', e.g. v0.1.0 (got: $TAG)"
fi

# build and push images
"${SCRIPTDIR}"/make-docker-images.sh

# update yaml
"${SCRIPTDIR}"/make-release-artifacts.sh

# create git release / push to new branch
git checkout -b "release/${TAG}"
git add "${SCRIPTDIR}/../release/"
git add "${SCRIPTDIR}/../kustomize/base/"
git commit --allow-empty -m "Release $TAG"
log "Pushing k8s manifests to release/${TAG}..."
git tag "$TAG"
git push --set-upstream origin "release/${TAG}"
git push --tags

log "Successfully tagged release $TAG."
