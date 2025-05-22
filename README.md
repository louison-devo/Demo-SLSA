# Demo SLSA

## Qu'est-ce que le framework SLSA ?

---

## Fonctionnement workflow

---

### Github action publish-image

---

La GitHub action est une composite action. Tout comme un reusable workflow, elle est réutilisable dans plusieurs workflows.<br>
Elle prend des paramètres d'entrées fixes et obligatoires(la **version de go**, le **token GitHub qui possède les permissions**, le **mot de passe de la registry** et le **nom de la container registry** ).<br>
Elle fournit en sortie le nom de l'image de conteneur et son digest générés par la GitHub action.<br>
Elle exécute plusieurs actions succesives:

1. **Installer Go** sur le runner de build
2. S'identifier à la Container registry (dans cette démo c'est la GitHub Registry)
3. **Installer Ko build** qui est l'outil permettant de builder et publier des images de container
4. **Installer Cosign** qui est l'outil édité par **Sigstore** permettant de signer des images de conteneurs
5. Utiliser la **Github action Goreleaser**. Goreleaser est un outil permettant de faire des **builds/releases automatisés** pour des projets en Golang principalement. Il est paramétré dans un fichier goreleaser.yml qui respecte une certaine [syntaxe](https://goreleaser.com/customization/).La publication d'images de conteneurs avec Ko est configurée dans ce fichier. Elle fournit en sortie un fichier d'archives.
6. Un script bash customs va permettre de trouver l'image et le digest de l'image de conteneurs générée
7. L'outil Cosign va permettre de signer l'image de conteneur et attaché directement la signature à l'image dans la registry

### GitHub action Verify Attestation
---

Une composite action est aussi utilisée pour exécuter le processus de verification automatisé de l'attestation de provenance au format SLSA. Il faut savoir qu'il est tout à fait possible de faire cette vérification à la main mais dans le but de respecter au maximum les critères du **niveau SLSA 3**, la vérification a été automatisée dans cette démo.<br>
Les différentes étapes de cette GitHub action sont:

1. Installation et configuration de Go
2. L'installation de l'outil de vérification **slsa-verifier** qui nécessite Go
3. Un script bash exécutant une commande avec slsa-verifier pour vérifier la conformité de l'attestation de provenance.

## Analyse de l'attestation de provenance

--- 

Lorsque ce Workflow est lancé, que tous les jobs sont réalisés et validés une image de conteneurs ainsi qu'une signature et une attestation sont stockées au niveau de la registry de l'image associée. L'outil de vérification permet de vérifier automatiquement la conformité de cette attestation mais on peut vouloir regarder les détails du build. Pour ce faire, il existe une méthode. <br>

On peut éxécuter la commande suivante pour enregistrer l'attestation de provenance:<br> 

`cosign download attestation ghcr.io/louison-devo/demo-slsa:[tag] | jq .payload -r | base64 --decode | jq > demo-slsa.att.json`
<br>

Ou bien si l'outil cosign n'est pas installé sur votre machine:
<br> 

`podman/docker run --rm -it gcr.io/projectsigstore/cosign:v2.5.0 download attestation ghcr.io/louison-devo/dem
o-slsa:[tag] | jq .payload -r | base64 --decode | jq > demo-slsa.attest.json`

## Lancement du fichier go

`docker run -d -p 8084:8084 --name demo-slsa ghcr.io/louison-devo/demo-slsa`
