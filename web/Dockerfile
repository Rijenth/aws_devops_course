FROM node:23.0.0

WORKDIR /app

COPY package.json pnpm-lock.yaml* ./

RUN npm install -g pnpm && pnpm install

COPY . .

EXPOSE 5173

CMD ["pnpm", "run", "dev", "--host"]
