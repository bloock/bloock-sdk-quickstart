DIR=$2;
MAVEN_OPTS=-Xss2M mvn -q compile assembly:single -Dname=$3 -Dpackage=com.bloock.quickstart.samples.${DIR//\//.}.$3;
