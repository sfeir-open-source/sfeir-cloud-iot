import React from 'react';
import { Bicycle } from 'components/bicycle/Bicycle';
import { Road } from 'components/road/Road';
import 'App.scss';
import { Footer } from 'components/footer/Footer';
import { useQuery } from 'react-query';
import { api } from 'http/http';
import { AxiosResponse } from 'axios';
import { round, sortBy } from 'lodash';

function App() {
  const { data: speed } = useQuery('speed', () => api.get('/speed').then((response: AxiosResponse<{ speed: number }>) => response.data.speed), { refetchInterval: 5100 });

  const { data: distance } = useQuery('distance', () => api.get('/distances', {
    params: {
      start: '2021-11-15',
      end: '2021-11-16'
    }
  }).then((response: AxiosResponse<{ distance: number }>) => response.data.distance), { refetchInterval: 5100 });

  const { data: rpm } = useQuery('rpm', () => api.get('').then((response: AxiosResponse<Record<string, { rpm: number }>>) => {
    const lastData: { rpm: number } | undefined = sortBy(response.data, (data, date) => new Date(date)).pop();
    return lastData?.rpm
  }), { refetchInterval: 5100 });

  return (
    <div className="app">
      <div className="info">
        <span className="speed">Speed: {round(speed || 0, 2)}km/h</span>
        <span className="distance">Distance: {round(distance || 0, 2)}km</span>
      </div>
      <div className="animation">
        <Road animationDuration={1}/>
        <Bicycle rpm={round(rpm || 0)}/>
      </div>
      <Footer/>
    </div>
  );
}

export default App;
