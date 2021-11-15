import 'components/bicycle/Bicycle.scss';
import { PedalLeft } from 'components/bicycle/Pedal/PedalLeft';
import { CrankArmLeft } from 'components/bicycle/CrankArm/CrankArmLeft';
import { WheelFront } from 'components/bicycle/Wheel/WheelFront';
import { SeatPost } from 'components/bicycle/Seat/SeatPost';
import { WheelRear } from 'components/bicycle/Wheel/WheelRear';
import { Brakes } from 'components/bicycle/Brakes';
import { Frame } from 'components/bicycle/Frame/Frame';
import { Seat } from 'components/bicycle/Seat/Seat';
import { PedalRight } from 'components/bicycle/Pedal/PedalRight';
import { CrankArmRight } from 'components/bicycle/CrankArm/CrankArmRight';
import { SeatPostTop } from 'components/bicycle/Seat/SeatPostTop';
import { FenderRearPost } from 'components/bicycle/Fender/FenderRearPost';
import { Caps } from 'components/bicycle/Caps';
import { HandleBars } from 'components/bicycle/HandleBars';
import { ChainTop } from 'components/bicycle/Chain/ChainTop';
import { ChainBottom } from 'components/bicycle/Chain/ChainBottom';
import { ChainRings } from 'components/bicycle/Chain/ChainRings';
import { Welding } from 'components/bicycle/Welding';
import { Cogset } from 'components/bicycle/Cogset/Cogset';
import { CogsetTop } from 'components/bicycle/Cogset/CogsetTop';
import { FrameSticker } from 'components/bicycle/Frame/FrameSticker';
import { FenderFront } from 'components/bicycle/Fender/FenderFront';
import { FenderRear } from 'components/bicycle/Fender/FenderRear';
import { Stem } from 'components/bicycle/Stem';
import { Stays } from 'components/bicycle/Stays';

interface BicycleProps {
  rpm: number;
}

export const Bicycle: React.FC<BicycleProps> = (props) => {
  const wheelAnimationDuration = 60 / props.rpm;
  const pedalAnimationDuration = wheelAnimationDuration * 2.6;
  const chainAnimationDuration = wheelAnimationDuration / 4;

  return (
    <svg className="bicycle" viewBox="0 0 588.35 360.02">
      <PedalLeft animationDuration={pedalAnimationDuration}/>
      <CrankArmLeft animationDuration={pedalAnimationDuration}/>
      <FenderFront/>
      <FenderRear/>
      <WheelFront animationDuration={wheelAnimationDuration}/>
      <SeatPost/>
      <WheelRear animationDuration={wheelAnimationDuration}/>
      <Cogset/>
      <Stem/>
      <Brakes/>
      <Frame/>
      <Seat/>
      <SeatPostTop/>
      <ChainTop animationDuration={chainAnimationDuration}/>
      <Stays/>
      <CogsetTop animationDuration={pedalAnimationDuration}/>
      <FenderRearPost/>
      <FrameSticker/>
      <Welding/>
      <Caps/>
      <HandleBars/>
      <ChainBottom animationDuration={chainAnimationDuration}/>
      <ChainRings animationDuration={pedalAnimationDuration}/>
      <CrankArmRight animationDuration={pedalAnimationDuration}/>
      <PedalRight animationDuration={pedalAnimationDuration}/>
    </svg>
  )
}