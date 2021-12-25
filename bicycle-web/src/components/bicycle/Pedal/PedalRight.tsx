interface PedalRightProps {
  animationDuration: number;
}

export const PedalRight: React.FC<PedalRightProps> = (props) => (
  <g className="pedal-right" style={{animationDuration: `${props.animationDuration}s`}}>
    <g className="pedal-right-inner" style={{animationDuration: `${props.animationDuration}s`}}>
      <rect className="stroke--round stroke--blue stroke--3px fill--yellow-light" x="296.94" y="305.36"
            width="44.03" height="10.5" transform="translate(-1.86 1.51) rotate(-0.36)"/>
      <rect className="stroke--round stroke--3px fill--yellow-dark stroke--blue" x="304.57" y="305.36"
            width="28.76" height="10.5" transform="translate(-1.86 1.51) rotate(-0.36)"/>
      <circle className="stroke--round stroke--3px fill--lavender stroke--blue" cx="318.95" cy="310.76"
              r="4.92"
              transform="translate(-1.86 1.51) rotate(-0.36)"/>
    </g>
  </g>
)