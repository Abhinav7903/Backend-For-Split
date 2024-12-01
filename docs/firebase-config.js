// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCb91proZ9tKMibxuKLdYAKX-NTI-7yuME",
  authDomain: "split-14044.firebaseapp.com",
  projectId: "split-14044",
  storageBucket: "split-14044.firebasestorage.app",
  messagingSenderId: "201453641673",
  appId: "1:201453641673:web:a2db42e24dd51820125218",
  measurementId: "G-475L0G850C"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
