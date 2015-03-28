import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0
import QtQuick.Dialogs 1.0

ApplicationWindow {
    visible: true
    title: "QML Bing Wallpaper"

    height: 576
    width: 1024
    minimumHeight: 576
    minimumWidth: 1024
    maximumHeight: 576
    maximumWidth: 1024

    //opacity: 0.9

    Image {
        source: imgURL
        smooth: true
        fillMode: Image.PreserveAspectFit
        height: parent.height
        width: parent.width
        anchors.centerIn: parent
    }

    Rectangle {
        height: 50
        width: parent.width
        y: parent.height - height
        color: "#AA000000"
        Text{
            text: descText
            font.family: "思源黑体", "Arial"
            font.pointSize: 24
            color: "white"
            anchors.centerIn: parent
        }
        Button{
            visible: false
            text: "fresh"
            onClicked:{
                buttonC();
            }
            
        }
    }
    property string imgURL
    property string descText
    signal buttonC
}
