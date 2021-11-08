import React from 'react';
import { Bicycle } from 'components/bicycle/Bicycle';
import { Road } from 'components/road/Road';
import 'App.scss';
import { Footer } from 'components/footer/Footer';

function App() {
  return (
    <div className="app">
      <div className="info">
        <span className="speed">Speed: 10km/h</span>
        <span className="distance">Distance: 100km</span>
      </div>
      <div className="animation">
        <Road/>
        <Bicycle/>
      </div>
      <Footer/>
    </div>
  );
}

export default App;
