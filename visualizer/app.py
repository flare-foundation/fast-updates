import argparse
from collections.abc import Hashable
from pathlib import Path
from typing import Any

import dash_bootstrap_components as dbc
import plotly.graph_objects as go
from dash import Dash, Input, Output, callback, dash_table, dcc, html
from plotly.subplots import make_subplots

from helpers import (
    get_log_paths,
    get_price_feeds,
    get_providers_status,
)

# Parse command line arguments
parser = argparse.ArgumentParser(
    prog="FastUpdatesVisualizer", description="Visualize Fast Updates live"
)
parser.add_argument(
    "-l",
    "--logpath",
    dest="logpath",
    default=Path("..", "logs"),
    type=Path,
    help="Path to the logs folder",
)
parser.add_argument(
    "-p", "--port", dest="port", type=int, default=8051, help="Port to run the app on"
)
args = parser.parse_args()
LOGPATH = args.logpath
PORT = args.port


@callback(
    [
        Output("live-update-table", "data"),
        Output("live-update-graph-feeds", "figure"),
        Output("live-update-graph-deviation", "figure"),
        Output("live-update-graph-status", "figure"),
    ],
    [
        Input("interval-component", "n_intervals"),
        Input("range-dropdown", "value"),
        Input("sd-slider", "value"),
    ],
)
def update(
    _, range_dropdown: str, sd_slider: float
) -> tuple[list[dict[Hashable, Any]], go.Figure, go.Figure, go.Figure]:
    """
    Update the visualizer with new data based on the selected range and
    standard deviation slider.

    Parameters:
        _: Placeholder parameter for the first positional argument.
        range_dropdown (str): The selected range from the dropdown menu.
        sd_slider (float): The value of the standard deviation slider.

    Returns:
        tuple[list[dict[Hashable, Any]], go.Figure, go.Figure]: A tuple
        containing the updated table, feeds figure, and status figure.
    """
    daemon_logpath, provider_logpaths = get_log_paths(LOGPATH)

    # Parse logfile and create table and graph
    df = get_price_feeds(daemon_logpath)
    num_feeds = len(df.columns) // 2

    # Only show the last 50 blocks
    if range_dropdown == "New":
        df = df.tail(300)

    # Maximum 2 columns for graphs
    if num_feeds < 3:
        num_rows, num_cols = 1, num_feeds
    else:
        num_rows, num_cols = num_feeds // 2, 2

    # Create DataTable
    table = df.iloc[::-1].reset_index().to_dict("records")

    # Create Graph for Price Feeds
    feeds_fig = make_subplots(
        rows=num_rows,
        cols=num_cols,
        subplot_titles=[f"Feed {i}" for i in range(num_feeds)],
    )
    for idx in range(num_feeds):
        feeds_fig.add_trace(
            {
                "x": df.index,
                "y": df[f"FastUpdateFeed_{idx}"],
                "name": f"Fast Update (Feed {idx})",
            },
            row=(idx // 2) + 1,
            col=(idx % 2) + 1,
        )
        feeds_fig.add_trace(
            {
                "x": df.index,
                "y": df[f"ActualFeed_{idx}"],
                "name": f"Actual (Feed {idx})",
            },
            row=(idx // 2) + 1,
            col=(idx % 2) + 1,
        )

        # Add standard dev bands
        if sd_slider == 0:
            continue
        std = df[f"ActualFeed_{idx}"].std() * sd_slider
        feeds_fig.add_trace(
            go.Scatter(
                name="SD Upper",
                x=df.index,
                y=df[f"ActualFeed_{idx}"] + std,
                mode="lines",
                marker={"color": "#444"},
                line={"width": 0},
                showlegend=False,
            )
        )
        feeds_fig.add_trace(
            go.Scatter(
                name="SD Lower",
                x=df.index,
                y=df[f"ActualFeed_{idx}"] - std,
                marker={"color": "#444"},
                line={"width": 0},
                mode="lines",
                fillcolor="rgba(68, 68, 68, 0.3)",
                fill="tonexty",
                showlegend=False,
            )
        )
    feeds_fig.update_layout(
        xaxis_title="Block Number",
        yaxis_title="Price",
        hovermode="x",
    )

    # Create Graph for Price Deviation
    deviation_fig = make_subplots(
        rows=1,
        cols=1,
    )
    for idx in range(num_feeds):
        y = (
            (df[f"FastUpdateFeed_{idx}"] - df[f"ActualFeed_{idx}"]).abs()
            / df[f"ActualFeed_{idx}"]
            * 1e4
        )
        deviation_fig.add_trace(
            {
                "x": df.index,
                "y": y,
                "name": f"Feed {idx} (med={y.median():.1f} bps, avg={y.mean():.1f} bps)",
            },
            row=1,
            col=1,
        )
    deviation_fig.update_layout(
        xaxis_title="Block Number",
        yaxis_title="Deviation (basis points)",
        hovermode="x",
        showlegend=True,
    )

    # Create Graph for Provider Status
    status_dict = get_providers_status(provider_logpaths)
    status_fig = make_subplots(rows=1, cols=1)
    for provider_id, status in status_dict.items():
        pct_failures = sum(status.values()) / len(status) if len(status) != 0 else 0

        status_fig.add_trace(
            {
                "name": f"Provider {provider_id} ({pct_failures:.0%} failures)",
                "x": list(status.keys()),
                "y": list(status.values()),
                "mode": "markers",
            },
            row=1,
            col=1,
        )
    status_fig.update_layout(
        xaxis_title="Block Number",
        yaxis_title="Failures",
        showlegend=True,
        hovermode="x",
    )
    return table, feeds_fig, deviation_fig, status_fig


# Initialize Dash app
app = Dash(__name__, external_stylesheets=[dbc.themes.BOOTSTRAP])
app.title = "Fast Updates Visualizer"
app.layout = dbc.Container(
    [
        html.H1("Fast Updates Visualizer"),
        html.Hr(),
        html.H2("Price Feeds"),
        dbc.Row(
            [
                dbc.Col(
                    [
                        dbc.Label("Range"),
                        dcc.Dropdown(
                            id="range-dropdown", options=["New", "All"], value="All"
                        ),
                    ]
                ),
                dbc.Col(
                    [
                        dbc.Label("Standard deviation"),
                        dcc.Slider(
                            id="sd-slider",
                            min=0,
                            max=2,
                            step=0.25,
                            value=1,
                            tooltip={"template": "SD={value}"},
                        ),
                    ]
                ),
            ],
            align="center",
        ),
        dcc.Graph(
            id="live-update-graph-feeds", style={"width": "100%", "height": "80vh"}
        ),
        dash_table.DataTable(id="live-update-table", page_size=10),
        html.Hr(),
        html.H2("Price Deviations"),
        dcc.Graph(
            id="live-update-graph-deviation", style={"width": "100%", "height": "80vh"}
        ),
        html.Hr(),
        html.H2("Provider Errors"),
        dcc.Graph(id="live-update-graph-status"),
        dcc.Interval(id="interval-component", interval=3000, n_intervals=0),
    ],
    fluid=True,
)

# Run the app
if __name__ == "__main__":
    app.run(port=PORT)
