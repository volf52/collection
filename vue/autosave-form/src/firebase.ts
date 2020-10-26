import firebase from "firebase/app";
import "firebase/firestore";

const firebaseApp = firebase.initializeApp({
  apiKey: "AIzaSyA7fCr4cl0LNtVFJ008MunRmzMPkJ3eapo",
  authDomain: "learning-c5530.firebaseapp.com",
  databaseURL: "https://learning-c5530.firebaseio.com",
  projectId: "learning-c5530",
  storageBucket: "learning-c5530.appspot.com",
  messagingSenderId: "299129820827",
  appId: "1:299129820827:web:1b35f8df6c6415e50f183e",
});

const db = firebaseApp.firestore();

export { db };
