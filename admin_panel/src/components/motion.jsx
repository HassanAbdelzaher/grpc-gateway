import { useSpring, animated } from 'react-spring'

function Motion() {
    const props = useSpring({
        config: { duration: 3000 },
        to: { number: 100 },
        from: { number: 1 }
    })
    return <animated.h1 >{props.number.interpolate(x => x.toFixed(0))}</animated.h1>
}

export default Motion;