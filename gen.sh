# Generator script for opensolar like platforms

# Find and replace YaleOpenLab/opensolar with the name desired - orgName / platform_name
echo "Enter Organization Name: "
read orgName
echo "Enter Platform Name:"
read repoName

if [ -z "$orgName" ] ; then
  echo "Empty string passed for org name, quitting"
  exit
fi

if [ -z "$repoName" ] ; then
  echo "Empty string passed for repo name, quitting"
  exit
fi

find . -name '*.go' -print0 | xargs -0 sed -i "" "s,github.com/YaleOpenLab/opensolar,github.com/$orgName/$repoName,g"
