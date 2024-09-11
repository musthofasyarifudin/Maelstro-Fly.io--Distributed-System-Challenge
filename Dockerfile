FROM ubuntu:latest

# Install Maelstrom
RUN apt-get update
RUN apt install openjdk-17-jdk graphviz gnuplot wget ruby-full git -y

# Set working directory
WORKDIR /challenge

RUN wget https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2
RUN tar -xvjf /challenge/maelstrom.tar.bz2 .

# Expose Maelstrom portr
EXPOSE 8080

RUN ./maelstrom test -w echo --bin demo/ruby/echo.rb --time-limit 5

RUN ./maelstrom serve
# Run Maelstrom with challenge script
#CMD ["maelstrom", "test", "challenge.clj"]
CMD ["tail", "-f", "/dev/null"]
