interface ChainTopProps {
  animationDuration: number
}

export const ChainTop: React.FC<ChainTopProps> = (props) => (
  <g className="chain-top">
    <line className="stroke--round fill--teal-light stroke--2px stroke--blue" x1="121.07" y1="222.52"
          x2="269.92"
          y2="239"/>
    <line className="stroke--round stroke--dasharray stroke--3px fill--teal-light stroke--blue chain-move"
          x1="121.07" y1="222.52" x2="269.92" y2="239"
          style={{animationDuration: `${props.animationDuration}s`}}
    />
  </g>
)