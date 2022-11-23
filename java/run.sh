DIR=$2;
mvn -q compile assembly:single -Dname=$3 -Dpackage=com.bloock.quickstart.samples.${DIR//\//.}.$3;