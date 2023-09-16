import { MantineProvider, createStyles, rem, Card, Container, Group, Badge, SimpleGrid, Text, AppShell, Space } from "@mantine/core"
import { useEffect, useState } from "react"
import { Response, Service } from "./types"

const useStyles = createStyles((theme) => ({
  title: {
    fontSize: rem(34),
    fontWeight: 900,

    [theme.fn.smallerThan('sm')]: {
      fontSize: rem(24),
    },
  },

  description: {
    maxWidth: 600,
    margin: 'auto',

    '&::after': {
      content: '""',
      display: 'block',
      backgroundColor: theme.fn.primaryColor(),
      width: rem(45),
      height: rem(2),
      marginTop: theme.spacing.sm,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  },

  card: {
    border: `${rem(1)} solid ${theme.colors.dark[5]}`,
    backgroundColor: theme.colors.dark[7],
    color: theme.colors.dark[0],
    '&:hover': {
      borderColor: theme.colors.dark[0],
    }
  },

  cardTitle: {
    '&::after': {
      content: '""',
      display: 'block',
      backgroundColor: theme.fn.primaryColor(),
      width: rem(45),
      height: rem(2),
      marginTop: theme.spacing.sm,
    },
  },
}));

export function FeaturesCards({ services }: { services: Service[] }) {
  const { classes } = useStyles();
  const features = services.map((service) => {
    if (service.ports.length === 0 || service.ip === "None") {
      return null
    }
    if (service.ports.length === 1) {
      return (
        <Card key={service.name} shadow="md" radius="md" className={classes.card} padding="xl" component="a" href={`http://${service.ip}:${service.ports[0].port}`}>
          <Text fz="lg" fw={500} className={classes.cardTitle} mt="md">
            {service.name}
          </Text>
          <Text fz="sm" c="dimmed" mt="sm">
            {"Service type:" + service.type}
            <Space />
            {`URL: ${service.ip}:${service.ports[0].port}`}
            <Space />
            {"Namespace: "+service.namespace}
          </Text>
        </Card>
      )
    }
    if (service.ports.length > 1) {
      service.ports.map((port) => {
        return (
          <Card key={`http://${service.ip}:${port.port}`} shadow="md" radius="md" className={classes.card} padding="xl" component="a" href={`http://${service.ip}:${port.port}`}>
            <Text fz="lg" fw={500} className={classes.cardTitle} mt="md">
              {service.name}
            </Text>
            <Text fz="sm" c="dimmed" mt="sm">
              {"Service type:" + service.type}
              <Space />
              {`URL: ${service.ip}:${port.port}`}
              <Space />
              {"Namespace: "+service.namespace}
            </Text>
          </Card>
        )
      })
    }
  });

  return (
    <Container size="lg" py="xl">
      <Group position="center">
        <Badge variant="filled" size="lg">
          XHOMER
        </Badge>
      </Group>

      <SimpleGrid cols={3} spacing="xl" mt={50} breakpoints={[{ maxWidth: 'md', cols: 1 }]}>
        {features}
      </SimpleGrid>
    </Container>
  );
}
function App() {

  const [services, setServices] = useState<Service[]>([])


  useEffect(() => {
    fetch('/api/v1/services')
      .then(response => response.json())
      .then((data: Response) => {
        setServices(data.services.reverse())
      })
  }, [])

  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <AppShell
        styles={(theme) => ({
          main: { backgroundColor: theme.colors.dark[8] },
        })}
      >
        <Container size="lg" py="xl">
          <FeaturesCards services={services} />
        </Container>
      </AppShell>
    </MantineProvider>
  )
}

export default App
