interface CrankArmRightProps {
  animationDuration: number;
}

export const CrankArmRight: React.FC<CrankArmRightProps> = (props) => (
  <g className="crankarm-right" style={{animationDuration: `${props.animationDuration}s`}}>
    <path className="stroke--round stroke--blue stroke--3px fill--yellow-light"
          d="M274.86,263.24a8.71,8.71,0,0,1,2,5.17c0.1,2.66,45.48,41.23,45.48,41.23l-3.9,3.88s-49.23-35.94-50.8-36A8.75,8.75,0,1,1,274.86,263.24Z"
          transform="translate(0.07 -0.48)"/>
    <circle className="stroke--round fill--lavender stroke--2px stroke--blue" cx="268.17" cy="267.95"
            r="2.33"/>
  </g>
)