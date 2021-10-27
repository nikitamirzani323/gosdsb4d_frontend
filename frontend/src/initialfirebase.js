import firebase from 'firebase/app';

function InitialFirebase(){
    var config = {
        apiKey: "AIzaSyCnjwV66P7jDLx5A0Hlh7CHKoZ2tg9jmMY",
        authDomain: "united-rope-233010.firebaseapp.com",
        databaseURL: "https://united-rope-233010-default-rtdb.asia-southeast1.firebasedatabase.app",
        projectId: "united-rope-233010",
        storageBucket: "united-rope-233010.appspot.com",
        messagingSenderId: "994050756260",
        appId: "1:994050756260:web:4dee40c4ca0c34a1842031"
    };
    firebase.initializeApp(config);
}

export default InitialFirebase