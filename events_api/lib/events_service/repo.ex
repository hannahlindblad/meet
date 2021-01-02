defmodule EventsService.Repo do
  use Ecto.Repo,
    otp_app: :events_service,
    adapter: Ecto.Adapters.Postgres
end
