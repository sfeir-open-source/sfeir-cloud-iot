import 'components/road/Road.scss';

interface RoadProps {
  animationDuration: number;
}

export const Road: React.FC<RoadProps> = (props) => {
  return (
    <div className="road">
      <div className="road-animation">
        <div className="road-lines" style={{animationDuration: `${props.animationDuration}s`}}/>
      </div>
    </div>
  )
}