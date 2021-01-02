# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
use Mix.Config

config :events_service,
  ecto_repos: [EventsService.Repo]

# Configures the endpoint
config :events_service, EventsServiceWeb.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "B2iGRP6OLuVgnnaSd/SpPO/+pFQozeLTdlHsBzJiI+pNcBHZL5+RbE6NyTo+JDJ1",
  render_errors: [view: EventsServiceWeb.ErrorView, accepts: ~w(json), layout: false],
  pubsub_server: EventsService.PubSub,
  live_view: [signing_salt: "qc8bP6Yx"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env()}.exs"
