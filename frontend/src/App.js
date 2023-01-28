import logo from './logo.svg';
import React, { useState } from 'react'
import './App.css';
import LoginForm from './Components/LoginForm/LoginForm';

function App() {
  const [ message, setMessage ] = useState(0);

  const ws = new WebSocket("ws://" + window.location.hostname + ":8080/ws");

  ws.onopen = () => {
    setMessage("Hello from opening !!")
  }
  return (
    <div className="App">
      <LoginForm></LoginForm>
    </div>
  );
}

export default App;
