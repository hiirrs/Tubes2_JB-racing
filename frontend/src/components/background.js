import React from 'react';
import { useSpring, animated } from 'react-spring';

const AnimatedBackground = () => {
    const styles = useSpring({
        from: { backgroundColor: 'red' },
        to: async next => {
            while (true) {
                await next({ backgroundColor: 'blue' });
                await next({ backgroundColor: 'green' });
                await next({ backgroundColor: 'red' });
            }
        },
        config: { duration: 5000 } // 5 seconds
    });

    return (
        <animated.div className="animated-background" style={styles}>
            {/* Your content goes here */}
        </animated.div>
    );
}

export default AnimatedBackground;
