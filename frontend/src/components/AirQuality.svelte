<script>
import { getContext } from 'svelte';

let current;
let label, lastSeen, pm2_5, pm2_5_10min_avg;

function refresh() {
    window.backend.NewAirQuality().then((result) => {
        let outside = result.results.filter((x) => x.DEVICE_LOCATIONTYPE === "outside");
        current = outside.shift()
        console.log(current);
        if (current != void 0) {
            label = current.Label;
            lastSeen = new Date(current.LastSeen*1000);
            pm2_5 = current.PM2_5Value;
            pm2_5_10min_avg = JSON.parse(current.Stats).v1;
        }
    });
}

function colorify(pm) {
    if ( pm <= 12 ) { return "green"; }
    if ( pm <= 35.4 ) { return "yellow"; }
    if ( pm <= 55.4 ) { return "orange"; }
    if ( pm <= 150.4 ) { return "red"; }
    if ( pm <= 250.4 ) { return "purple"; }
    return "maroon";
}


refresh();

        
</script>
<h1>HMB Air Quality</h1>
<main>
    {#if current != null}
    <table>
        <tr>
            <th> Where </th>
            <th> Status </th>
            <th> PM2.5 Level </th>
        </tr>
        <tr>
            <td> {label} </td>
            <td class={colorify(pm2_5_10min_avg)}>  </td>
            <td> {pm2_5_10min_avg} </td>
        </tr>
    </table>
    <p> last update at <br/> {lastSeen.toLocaleString()} </p>
    {/if}

    <p><button on:click={refresh}>refresh</button></p>

    <ul>
        <li> <a href="https://ww2.arb.ca.gov/resources/inhalable-particulate-matter-and-health">Inhalable Particulate Matter and Health</a> </li>
    </ul>
</main>

<style>
.green {
    background-color: green;
}
.yellow {
    background-color: yellow;
}
.orange {
    background-color: orange;
}
.red {
    background-color: red;
}
.purple {
    background-color: purple;
}
.maroon {
    background-color: maroon;
}
</style>
