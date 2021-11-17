import React from 'react';
import { Bicycle } from 'components/bicycle/Bicycle';
import { Road } from 'components/road/Road';
import 'App.scss';
import { Footer } from 'components/footer/Footer';
import { useQuery } from 'react-query';
import { api } from 'http/http';
import { AxiosResponse } from 'axios';
import { round, sortBy } from 'lodash';

const today = new Date();
const tomorrow = new Date();
tomorrow.setDate(today.getDate() + 1);

const year = today.getFullYear();
const month = today.getMonth() + 1;

function App() {
  const { data: speed } = useQuery('speed', () => api.get('/speed').then((response: AxiosResponse<{ speed: number }>) => response.data.speed), { refetchInterval: 5100 });

  const { data: distance } = useQuery('distance', () => api.get('/distances', {
    params: {
      start: `${year}-${month}-${today.getDate()}`,
      end: `${year}-${month}-${tomorrow.getDate()}`
    }
  }).then((response: AxiosResponse<{ distance: number }>) => response.data.distance), { refetchInterval: 5100 });

  const { data: rpm } = useQuery('rpm', () => api.get('').then((response: AxiosResponse<Record<string, { rpm: number }>>) => {
    const lastData: { rpm: number } | undefined = sortBy(response.data, (data, date) => new Date(date)).pop();
    return lastData?.rpm
  }), { refetchInterval: 5100 });

  return (
    <>
      <div className="background"/>
      <div className="app">
        <div className="info">
          <span className="speed">Vitesse: {round(speed || 0, 2)}km/h</span>
          <span className="distance">Distance: {round(distance || 0, 2)}km</span>
        </div>
        <div className="animation">
          <Road animationDuration={speed ? 3 / (speed || 1) : 0}/>
          <Bicycle rpm={speed ? round(rpm || 0) : 0}/>
        </div>
        <div className="info">
          <span className="objective">Objectif: 100km</span>
        </div>
        <Footer/>
      </div>
    </>
  );
}

export default App;
