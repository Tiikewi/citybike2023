import React from "react";

export const Map: React.FC<{lat: number, lng: number}> = ({lat, lng}) => {
    const ref = React.useRef<HTMLDivElement>(null);
    const [map, setMap] = React.useState<google.maps.Map>();

    React.useEffect(() => {
        const location: google.maps.LatLngLiteral = {lat, lng};

        if (ref.current && !map) {
            setMap(new window.google.maps.Map(ref.current, {center: location, zoom: 16}));
        }

        new google.maps.Marker({
            position: location,
            map: map
          });
    }, [ref, map, lat, lng]);

    return <div ref={ref} style={{width: "100%", height: "100%", margin: "8px"}} />
}


  